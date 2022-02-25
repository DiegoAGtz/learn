// let A = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [12, 14, 15, 16]];
//let A = [[0, 6], [5, 4]];
let A = [[112, 42, 83, 119], [56, 125, 56, 49], [15, 78, 101, 43], [62, 98, 114, 108]];

let matrix = (arr) => {

  let swap = (arr) => {
    let reducer = (a, b) => a+b;
    for (let i = 0; i < arr.length; i++) {
      let tmp1 = arr[i].slice(0, arr.length/2);
      let tmp2 = arr[i].slice(arr.length/2, arr.length)

      if (tmp1.reduce(reducer) < tmp2.reduce(reducer)) {
        arr[i] = [...tmp2.reverse(), ...tmp1.reverse()];
      }
    }
    return arr;
  }

  let transpose = (arr) => {
    let arrT = new Array(arr.length);
    for (let i = 0; i < arr.length; i++) {
      let row = new Array(arr.length);
      for (let j = 0; j < arr.length; j++) {
        row[j] = arr[j][i];
      }
      arrT[i] = row;
    }
    return arrT;
  }

  let res = 0;

  let arrsw = swap(arr);
  let arrT = transpose(arr);
  arrsw = swap(arrT);

  for (let i = 0; i < arr.length/2; i++) {
    for (let j = 0; j < arr.length/2; j++) {
      res += arrsw[i][j];
    }
  }
  console.log(res);
  printArr(arrsw);
  return res;
}


let printArr = (arr) => {
  for (let i = 0; i < arr.length; i++) {
    console.log(arr[i]);
  }
  console.log('');
}

matrix(A);
