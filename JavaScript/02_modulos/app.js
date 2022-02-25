// require('./datos');
// require('./carpeta/modulo');

const datos = require('./datos'); // Importa el módulo datos con el alias datos.
datos.log('¡Hola mundo!');
console.log(module);

// ******* Algunos módulos integrados *******
console.log(__filename);    // Variable que hace referencia al archivo local
console.log(__dirname);    // Variable que hace referencia al directorio local

const path = require('path');
const objPath = path.parse(__filename);
console.log(objPath);

// ´Módulo OS
const os = require('os');
var memoriaLibre = os.freemem();
var memoriaTotal = os.totalmem();

console.log(memoriaLibre);
console.log(memoriaTotal);

// Módulo eventos
const EventEmitter = require('events'); // Importa la clase eventos
const emiter = new EventEmitter(); // Declara una instancia de la clase eventos

// Regristrar el listener
emiter.on('mensajeLoger', function() {
    console.log('Listener llamado...');
}),
emiter.emit('mensajeLoger');
emiter.emit('mensajeLoger', {id: 1, url: 'http://test.com'});
