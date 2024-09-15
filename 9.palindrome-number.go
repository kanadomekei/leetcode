/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
/*
	数字を配列Aとして、取得する。
	逆向きの配列Bも作成する。
	配列Aと配列Bが同じかどうか判断する
	→計算量はO(n)だと思う。
 */
// @lc code=start
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    // 数字を配列Aとして取得
    var a []int
    for i := x; i > 0; i /= 10 {
        a = append(a, i%10)
    }

    // 逆向きの配列Bを作成
    b := make([]int, len(a))
    for i := 0; i < len(a); i++ {
        b[i] = a[len(a)-1-i]
    }

    // 配列Aと配列Bが同じかどうか判断
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}
// @lc code=end

