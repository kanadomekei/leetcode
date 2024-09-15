/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
/*
	文字列をを配列として取得する。
	symbolからvalueへのmapを用意する
	sum=0で初期化する・
	一文字ずつ変換して足していく。	
	→単調にすべて足すだけではだめらしい。
	以下のパターンに対応するために、右側から処理をする。
	Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

 */
// @lc code=start
func romanToInt(s string) int {
    // シンボルから値へのマップを用意する
    romanValues := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    
    sum := 0
    prevValue := 0
    
    // 文字列を右から左へ一文字ずつ処理する
    for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanValues[s[i]]
        
        // 現在の値が前の値以上の場合は加算、そうでない場合は減算
        if currentValue >= prevValue {
            sum += currentValue
        } else {
            sum -= currentValue
        }
        
        prevValue = currentValue
    }
    
    return sum
}
// @lc code=end

