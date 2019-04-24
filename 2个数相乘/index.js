function multipy(num1, num2) {
    var array1 = reverse(toArrayInt(num1));
    var array2 = reverse(toArrayInt(num2));
    var num = new Array()
    for(var i=0;i<array1.length+array2.length;i++){
        num[i]=0
    }
    for (var i=0;i<array1.length;i++){
         for(var j=0;j<array2.length;j++){
             num[i+j]+=array1[i]*array2[j]
         }
    }
    var c=0
    for(var i=0;i<num.length;i++){
        var t= num[i]+c
        num[i]= t%10
        c =parseInt(t/10)
    }

    num = reverse(num)
    var flag= false
    var s=""
    for (var i=0;i<num.length;i++){
        if(flag){
            s+=num[i].toString()
        }else {
            if(num[i]!=0){
                s+=num[i].toString()
                flag= true
            }
        }
    }
    console.log(s)
}

function toArrayInt(num) {
    var list = new Array()
    for (var i = 0; i < num.length; i++) {
        list[i] = parseInt(num[i])
    }
    return list
}

function reverse(list) {
    var i = 0;
    var j = list.length - 1;
    while (i < j) {
        var m = list[i];
        list[i] = list[j];
        list[j] = m;
        i++;
        j--;
    }
    return list
}
multipy('25','26');