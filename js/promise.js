// setTimeout(function() {
//     console.log(10)
// }, 10);
// setTimeout(function() {
//     console.log(1)
// }, 0);
// new Promise(function executor(resolve) {
//     console.log(2);
//     for( var i=0 ; i<10000 ; i++ ) {
//         i == 9999 && resolve();
//     }
//     console.log(3);
// }).then(function() {
//     console.log(4);
// });
// console.log(5);

//输出 2 3  5  4  1 10


let doSth = new Promise((resolve, reject) => {
        console.log('hello');
        throw new Error('悲剧了，又出 bug 了');
        reject({a:1});
});
doSth.then(() => {
    console.log('over');
},function (err) {
   console.log(err,"err")
}).catch(function(err){
    console.log(err,"err");
});