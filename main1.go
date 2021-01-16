func lengthOfLongestSubstring(s string) int {
    l:=len(s);
    if(l<2){
        return l;
    }
    // longest no dup string length
    longestInt := 0;

    var currentString []rune;
    givenString := []rune(s); 
    
    for _, c := range givenString {
        currentString = processString(currentString, c);
        l:=len(currentString);
        if(l>longestInt) {
            longestInt = l;
        }
        //longestInt = Max(longestInt, len(currentString));
	}
    return longestInt;
}
/*
func Max(x int, y int) int{
    if x < y {
        return y
    }
    return x
}*/

/*
func processString( st []rune, c rune) []rune {
    
    
    //var newString []rune;
    
    //pos := strpos(st, c);
    
    if(pos==-1){
        newString=append(st,c);
    }else{
        if(pos+1 == len(st)){ // last letter 
            newString = nil;  // delete all array
        }else{
            newString = st[pos+1:]; 
        }
        newString = append(newString,c); // add the new char
    }
    return newString;
}*/

// return the position of c in st; -1 if not found
// also return the rest of chars if found
func processString (st []rune, v rune) ( []rune) {
    for i, ch := range st {
        if (ch == v){
            if(i+1 == len(st)){
                return []rune{v};
            }else{
                return append(st[i+1:],v);
            }
            // return i;
        }
        i++;
    }
    return append(st, v);
}
