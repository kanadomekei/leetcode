/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

 /*
    素直に全探索したら解けそうだけど、O(n^2)になる
    go言語の場合、配列の中にある値が存在するか知るための処理の計算量はいくつだろう？
    →二部探索でO(logn)
    mapを使うのが良いらしい
 */
// @lc code=start
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, exists := numMap[complement]; exists {
            return []int{j, i}
        }
        numMap[num] = i
    }
    return nil
}
// @lc code=end
 /*
    https://leetcode.com/problems/two-sum/?source=vscode
    全探索の場合はメモリ効率が良い。計算量はO(n^2)

    ソートする場合のアルゴリズムもある。O(nlogn)
    1,配列をソートする（まだソートされていない場合）
    2,2つのポインタを配置します。1つは配列の先頭に、もう1つは配列の末尾に置きます。
    3,これら2つのポインタの要素の合計を計算します
    4,合計が目標値と等しければ、解決策が見つかったことになります
    5,合計が目標値より小さい場合は、左のポインタを右に移動します
    6,合計が目標より大きい場合は、右のポインタを左に移動します
    7,解決策が見つかるか、ポインタが一致するまで、手順3〜6を繰り返します。

    ソートされていない配列のインデックスが必要になるために、配列のコピーが必要

 */