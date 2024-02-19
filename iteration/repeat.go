package iteration

//  Repeat takes in a string and concatenates it for repeatCount times
func Repeat(character string, repeatCount int) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += character
    }
    return repeated 
}
