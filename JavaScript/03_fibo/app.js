const serie = require('./serie');
const process = require('process');

let argv = process.argv;
let cantidad = argv[2].split('=');
console.log(cantidad);

serie.crearSerie(cantidad)
    .then(mensaje => console.log(mensaje))
    .catch(mensaje => console.log(mensaje));