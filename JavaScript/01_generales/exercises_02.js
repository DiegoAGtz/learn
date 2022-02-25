const prompt = require('prompt-sync')();
console.log('Ejercicio 03');

let menores = 0;
let mayores = 0;
for (let i = 0; i < 10; i++) {
  let edad = prompt(`Ingresa la edad ${i+1}: `);
  (Number(edad) < 18) ? menores++ : mayores++;
}
console.log(`Hay ${menores} menores de edad y ${mayores} mayores de edad.`);
