const inq = require('inquirer');
const fig = require('figlet');
const run = async () => {

    fig.text('Shell TC',
        function (err,data) {
        if(err){
            return;
        }
        console.log(data);
        });



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




