func detectCapitalUse(word string) bool {
    var res bool = false
    
    //case 1
    if(IsUpper(word)){
        res = true
        return res
    }
    //case 2
    if(IsLower(word)){
        res = true
        return res
    }    
    //case 3
    
    result := []rune(word)
    firstCharacter := (result[0:1])[0]
    if unicode.IsUpper((firstCharacter)){
        res=true
        rem:= string(result[1:])
         for _,each := range(rem){
             if !unicode.IsLower(each){
                 res = false
                 break
             }
         }
        return res
    }
    return res
}

func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}