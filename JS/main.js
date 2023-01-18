import inq from 'inquirer';
import boxen from 'boxen';
import psList from 'ps-list';
import cp from 'child_process';
import { exec } from 'child_process'
import * as process from "process";
import { suspend, resume } from 'ntsuspend';
import chalk from 'chalk';
import fs from 'fs';


// variables
var cmdhistorytab = [];
var running = true;
var lastcommand = null;
var mainpath = process.cwd();
const actions = [
    { // La je fais un fonction dans l'array pour avoir des types
        name: 'lp',
        desc: 'lists the running processes on the machine'
    },
    {
        name: 'clear',
        desc: 'clears the console'
    },
    {
        name: 'exec (path-to-program)',
        desc: 'runs a program from PATH variables or direct path'
    },
    {
        name: 'bing (-k|-p|-c) (process_id)',
        desc: "-k kills select process -p pauses and -c resumes"
    },
    {
        name: 'dir',
        desc: 'displays the files and folders in the current directory'
    },
    {
        name: 'cd (directory)',
        desc: 'navigate through directories'
    },
    {
        name: 'exit',
        desc: 'exit the CLI. Ctrl+P also works'
    },
    {
        name: 'history',
        desc: 'Shows previous commands'
    }
    ]

// initiale clear + intro
console.clear();
console.log(chalk.yellow(boxen('Shell TC v1.3', {padding: 1})));
console.log(chalk.green("Run 'help' to see available commands "));

//fonctionenment
const run = async () => {

    exitCommand();
    mainpath = process.cwd();
    const cmd = await line();
    cmdhistorytab.push(cmd); // history of commands
    await action(cmd);
}

function line() {
    return inq.prompt(
        [
            {
                name: 'command',
                type: 'input',
                message: chalk.blue(mainpath + ' $ ')

            }
        ]);
}


//actions
async function action (cmd) {

    if(cmd.command === "lp"){

        console.log(chalk.green(boxen('Running processes')));

      let processes = (await psList());
      processes.sort(compare);
      processes.reverse(); // plus gros d'abord - souvent les plus utilisés? car OS et tout démarre avant

      console.log('\n');
      for(let i = 0; i < 30; i++){
          console.log(chalk.yellow(i+1)+'. '+ processes[i].name+' pid: ' + chalk.red(processes[i].pid) +"\n");
      }
    }

    else if(cmd.command === "clear"){
        console.clear();
    }

    else if(/^exec /.test(cmd.command)) { // REGARDER PATH VARIABLES !
        let prog = cmd.command.replace(/^exec /, "");
        //console.log(prog);

        // Execute the command
        await cp.exec(prog, (error) => {
            if (error) {
                console.error(`exec error: ${error}`);
            }
        });
    }

    else if(/^bing/.test(cmd.command)) {

        const com = cmd.command.match(/(-k|-p|-c) /);
        const processId = cmd.command.replace(/^bing (-k|-p|-c) /,"");

            if(com != null) {
                switch (com[0]) {
                    case '-k ': //Faire gaffe aux espaces dans les case
                        // Kill the process

                        try{
                            process.kill(processId);
                            console.log("Process killed:  " + processId);
                        }catch{
                            console.log("Process doesn't exist ! Use lp to see process list.")
                        }
                        break;
                    case '-p ':
                        // Pause the process
                        console.log("Process paused: " + processId);
                        if (process.platform === 'win32') {
                            suspend(processId);
                        }else{
                            try {
                                process.kill(processId, 'SIGSTOP');
                            }catch{
                                console.log("Process doesn't exist ! Use lp to see process list.")
                            }
                        }
                        break;
                    case '-c ':
                        // Resume the process
                        console.log("Process resumed: " + processId);
                        if (process.platform === 'win32') {
                            resume(processId);
                        }else{
                            process.kill(processId, 'SIGCONT');
                        }
                        break;

                }
            } else {
                console.error('Use bing -k|-p|-c (process_id) instead');
            }
        }

    else if(/^cd(..$| )/.test(cmd.command)) {

        let pathname = cmd.command.replace(/^cd ?/,"");

        if(process.platform !== 'win32'){
            if(pathname === '..'){
                pathname = '../';
            }
        }

        try{
            process.chdir(pathname);
        }catch (e) {
            console.log("Please enter a valid path !");
        }
    }

    else if ((cmd.command) === 'dir'){

        fs.readdir(mainpath, (err, files) => {
            console.log("\n");
            for(let i = 0 ; i < files.length ; i++) {
                if (/[.]/.test(files[i])) {
                    console.log(chalk.red(files[i]));
                } else {
                    console.log(chalk.yellowBright(files[i]));
                }
            }
        });
    }

    else if (cmd.command === 'help'){

        for(let i = 0 ; i < actions.length ; i++){
            console.log(chalk.whiteBright(actions[i].name) + " -- " + chalk.magenta(actions[i].desc));
        }
    }

    else if (cmd.command === 'exit') {
        process.exit();
        running = false;
    }

    else if (/^!/.test(cmd.command)) {
        var prog;
        if(process.platform === 'win32'){
            prog = 'start /min ' + cmd.command.replace(/^!/, "");
        } else {
            prog = cmd.command.replace(/^!/, "") + '&';
        }

        const child = await cp.exec(prog, { detached: true, stdio: 'ignore'}, async (error) => {
            if (error) {
                await console.error(`exec error: ${error}`);
            }
        });
        child.unref();
        child.stdout.on('data', data => {
            console.log(`stdout: ${data}`);
        });

        child.stderr.on('data', data => {
            console.error(`stderr: ${data}`);
        });

        child.on('close', code => {
            console.log(`child process exited with code ${code}`);
        });

    }

    else if(/^keep /.test(cmd.command)) {
        //garder prog
        let processId = cmd.command.replace(/^keep /, "");
        if(process.platform === 'win32'){
            cp(`start /B ${processId}`, (err, stdout, stderr) => {

                console.log(stdout);
            });
        } else {

            exec(`nohup disown ${processId}`, (err, stdout, stderr) => {

                console.log(stdout);
            });

        }


    }

    else if(cmd.command === "history") {

        for(let i = 0; i < cmdhistorytab.length; i++){
            console.log(chalk.yellow(i+1)+'. '+ cmdhistorytab[i].command +"\n");
        }
    }

    else {console.log("Unrecognized command:" + cmd.command)}
}

function exitCommand(){
    var stdin = process.stdin;

    stdin.setRawMode(true);
    stdin.resume();
    stdin.setEncoding('utf8');

    stdin.on('data', key => {
        // Ctrl+P
        if (key === '\u0010') {
            process.exit();
            running = false;
            console.log("bug");
        }
    });
}




//fonction pour trier les list processes par celui qui a le plus de mémoire comme ca on voit pas
//tous les process
//marche pas sur windows
function compare( a, b) {
    if ( a.memory < b.memory){
        return -1;
    }
    if ( a.memory > b.memory){
        return 1;
    }
    return 0;
}




async function main(){
    while(running) {
        await run();
    }
}

main();