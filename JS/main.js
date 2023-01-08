import inq from 'inquirer';
import boxen from 'boxen';
import path from 'path';
import psList from 'ps-list';
import cp from 'child_process';
import * as process from "process";
import { suspend, resume } from 'ntsuspend';




// constantes
var running = true;
const mainpath = path.resolve('main.js');

// initiale clear + intro
console.clear();
console.log(boxen('Shell TC v1.2', {padding: 1}));

// fonctions
const run = async () => {

    exitCommand();
    const cmd = await line();
    //console.log(com["command"]);
    await action(cmd);
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
      processes.sort(compare);
      processes.reverse(); // plus gros d'abord - souvent les plus utilisés? car OS et tout démarre avant

      console.log('\n');
      for(let i = 0; i < 30; i++){
          console.log(i+1+'. '+ processes[i].name+' pid: ' + processes[i].pid +"\n");
      }
    }

    else if(cmd.command === "clear"){
        console.clear();
    }

    else if(/^exec /.test(cmd.command)) { // REGARDER PATH VARIABLES !
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
                    case '-k ': //Faire gaffe aux espaces dans les case
                        // Kill the process
                        console.log("kill " + processId);

                            process.kill(processId);


                        break;
                    case '-p ':
                        // Pause the process
                        console.log("pause " + processId);
                        if (process.platform === 'win32') {
                            suspend(processId);
                        }else{
                            process.kill(processId, 'SIGSTOP');
                        }
                        break;
                    case '-c ':
                        // Resume the process
                        console.log("continue " + processId);
                        if (process.platform === 'win32') {
                            resume(processId);
                        }else{
                            process.kill(processId, 'SIGCONT');
                        }
                        break;
                    default:
                        console.error(`Invalid command: ${com}`);
                        break;
                }
            } else {
                console.error(`Command is null : ${com}`);
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


//fonction pour trier les list processes par celui qui a le plus de mémoire comme ca on voit pas
//tous les process
//marche pas sur windows
function compare( a, b ) {
    if ( a.memory < b.memory ){
        return 1;
    }
    if ( a.memory > b.memory ){
        return -1;
    }
    return 0;
}


async function main(){
    while(running) {
        await run();
    }
}

main();