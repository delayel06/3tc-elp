import inq from 'inquirer';
import boxen from 'boxen';
import path from 'path';
import psList from 'ps-list';
import cp from 'child_process';


// constantes
var running = true;
const mainpath = path.resolve('main.js');

// initiale clear + intro
console.clear();
console.log(boxen('Shell TC v1.1', {padding: 1}));

// fonctions
const run = async () => {

    exitCommand();
    const com = await line();
    //console.log(com["command"]);
    await action(com);
}

function line() {
    return inq.prompt(
        [
            {
                name: 'command',
                type: 'input',
                message: mainpath + ' $ '
            }
        ]);
}

async function action (cmd) {

    if(cmd.command === "lp"){

      console.log(boxen('Running processes'));

      let processes = (await psList());

      console.log('\n');
      for(let i = 0; i < processes.length ; i++){
          console.log(i+'. '+ processes[i].name+'\n');
      }
    }

    else if(cmd.command === "clear"){
        console.clear();
    }

    else if(/^exec /.test(cmd.command)) {
        let prog = cmd.command.replace(/^exec /, "");
        //console.log(prog);

        // Execute the command
        await cp.exec(prog, (error, stdout, stderr) => {
            if (error) {
                console.error(`exec error: ${error}`);
            }
        });
    }

    else if(/^bing /.test(cmd.command)) {

        const com = cmd.command.match(/(-k|-p|-c) /);
        const processId = cmd.command.replace(/^bing (-k|-p|-c) /,"");

            if(com != null) {
                switch (com[0]) {
                    case '-k ':
                        // Kill the process
                        console.log("kill " + processId);
                        break;
                    case '-p ':
                        // Pause the process
                        console.log("pause " + processId);
                        break;
                    case '-c ':
                        // Resume the process
                        console.log("continue " + processId);
                        break;
                    default:
                        console.error(`Invalid command: ${com}`);
                        break;
                }
            } else {
                console.error(`Invalid command: ${com}`);
            }
        }

    else {console.log(cmd.command)}
}

function exitCommand(){
    const stdin = process.stdin;

    stdin.setRawMode(true);
    stdin.resume();
    stdin.setEncoding('utf8');

    stdin.on('data', key => {
        // Ctrl+P
        if (key === '\u0010') {
            process.exit();
            running = false;
        }
    });
}

async function main(){
    while(running) {
        await run();
    }
}

main();