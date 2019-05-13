
function twosum(array,target) {
    var m ={};
    for (var i=0;i<array.length;i++){
        // if(typeof m[0-array[i]] !='undefined'){
        //    return [m[0-array[i]],i];
        // }
        if(0-array[i] in m){
            return [m[0-array[i]],i];
        }
        m[array[i]-target] = i
    }
    return []
}
console.log(twosum([1,2,3,4],3))