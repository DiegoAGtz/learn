var url = 'https://wikipedia.or';

function dato(mensaje) {
    console.log(mensaje)
}

module.exports.log = dato;  // Exporta la función dato con el alias log
// module.exports.url = url;   // Exporta la variable url con el mismo nombre

// Las exportaciones se vuelven públicas, para ser accedidas desde otros programas
console.log(module);