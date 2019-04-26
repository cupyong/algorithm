function callbackLongest(s) {
    if(s.length<2){
        return s
    }
    var maxStart =0
    var maxLen =0
    for (var i=0;i<s.length;i++){
        var left=i,right =i;
        while (left>=0&&right<=s.length&&s[left]==s[right]){
            if(maxLen<right-left+1){
                maxLen =right-left+1
                maxStart = left
            }
            left--
            right++
        }
        var left=i,right =i+1;
        while (left>=0&&right<=s.length&&s[left]==s[right]){
            if(maxLen<right-left+1){
                maxLen =right-left+1
                maxStart = left
            }
            left--
            right++
        }
    }
    return s.substring(maxStart,maxLen)
}
function test() {
    console.log(callbackLongest('adffdaasdasd'))
    console.log(callbackLongest("adasdf"))
}
test()