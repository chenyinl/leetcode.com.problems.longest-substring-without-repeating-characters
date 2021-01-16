func lengthOfLongestSubstring(s string) int {
    
    // initial the total length
    l := len(s);
    if(l<2){
        return l;
    }
    
    longestInt := 0;   // the answer;
    
    // point the pointer
    p_start := 0; // start of current string
    p_end   := 0; // end of current string
    
    //i:=1;
    for i:=1; i<l; i++ {
        
        for j:=p_start; j<=p_end; j++{
            if(s[j] == s[i]){ // there is a duplicate'
                //p_start = j+1;
                if(j+1>l){
                    p_start = j;
                }else{
                    p_start = j+1; // p_start is at the end
                    //givenString[j]='';
                }
                continue;
            }
        }
        if(p_end<l-1){
            p_end++;
        }
       // fmt.Printf("[%d:%d]", p_start, p_end);
        
        tempLen:= p_end - p_start+1;
        if(tempLen>longestInt){
            longestInt = tempLen;
        }
    }
    //fmt.Println();
    //givenString=[]rune("");
    s="";
    return longestInt;
}

