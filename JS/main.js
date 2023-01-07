import inq from 'inquirer';
import boxen from 'boxen';
import path from 'path';
import psList from 'ps-list';


// constantes
const running = true;
const mainpath = path.resolve('main.js');

// initiale clear + intro
console.clear();
console.log(boxen('Shell TC v1.1', {padding: 1}));

// fonctions
const run = async () => {

    const com = await line();
    console.log(com);
    action(com);
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
    if(cmd.command === "clear"){
        console.clear();
    }
}

async function main(){
    while(running) {
        await run();
    }
}

main();