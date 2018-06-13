/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
    res = LCA(root,p,q);
    return res.best;
};

var LCA = function(root, p, q) {
    var res = {best : 0, ok : false}
    if (isAncestor(root,p) && isAncestor(root, q)) {
        res.best = root;
        res.ok = true;

        var left = {best: 0, ok: false};
        if (root.left !== null ){
            left = LCA(root.left,p,q)
        }

        var right = {best: 0, ok: false};
        if(root.right !== null) {
            right = LCA(root.right,p,q)
        }

        if(left.ok) {
            res = left;
        }

        if(right.ok) {
            res = right;
        }
    }
    return res
}

var isAncestor = function(node, p) {
    if (node == null){
        return false;
    }
    if(node == p) {
        return true;
    }

    return isAncestor(node.left, p) || isAncestor(node.right, p);
}

var main = function() {
    var root = {
        val : 6,
        left: {
            val: 2,
            left: {
                val: 0,
                left: null,
                right: null
            },
            right: {
                val: 4,
                left: {
                    val: 3,
                    left: null,
                    right: null
                },
                right:{
                    val: 5,
                    left: null,
                    right: null
                }
            }
        },
        right:{
            val:8,
            left: {
                val: 7,
                left: null,
                right: null
            },
            right: {
                val: 9,
                left: null,
                right: null
            }
        }
    };

    var p = 2;
    var q = 8;

    console.log(lowestCommonAncestor(root,p,q))
}

main()