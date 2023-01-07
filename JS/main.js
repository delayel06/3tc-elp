const inq = require('inquirer');

const run = async () => {

    var a = inq.prompt(
        [
            {
            name: 'command',
            type: 'input',
            message: 'Entrez la commande: '
            }
        ]);

    a.then(function (answer) {
        console.log("entrée : " + answer.command);
    });
}

run();




