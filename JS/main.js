import inq from 'inquirer';
import boxen from 'boxen';
import path from 'path';


console.clear();
console.log(boxen('Shell TC v1.1', {padding: 1}));

var mainpath = path.resolve('main.js');



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

function action (cmd) {

    if(cmd.command === "test"){
        console.log("hurrah");

    }


}


run();




