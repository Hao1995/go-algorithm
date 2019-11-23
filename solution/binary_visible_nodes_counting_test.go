package solution

import (
	"testing"

	"github.com/Hao1995/algorithm/models"
)

func TestBinaryTreeVisibleNodesCounting(t *testing.T) {

	var bTree *models.Tree
	var ans int

	// (5, (3, (20, None, None), (21, None, None)), (10, (1, None, None), None)) >> 4 (5,20,21,10)
	bTree = &models.Tree{
		Val: 5,
		Right: &models.Tree{
			Val: 10,
			Right: &models.Tree{
				Val: 1,
			},
		},
		Left: &models.Tree{
			Val: 3,
			Right: &models.Tree{
				Val: 21,
			},
			Left: &models.Tree{
				Val: 20,
			},
		},
	}
	ans = 4
	if val := BinaryTreeVisibleNodesCounting(bTree); val != ans {
		t.Errorf("Tree=%v. Expect %v, but get %v", "(5, (3, (20, None, None), (21, None, None)), (10, (1, None, None), None)) >> 4 (5,20,21,10)", ans, val)
	}

	// (8, (2, (8, None, None), (7, None, None)), (6, None, None)) >> 2 (8,8)
	bTree = &models.Tree{
		Val: 8,
		Right: &models.Tree{
			Val: 3,
		},
		Left: &models.Tree{
			Val: 2,
			Right: &models.Tree{
				Val: 7,
			},
			Left: &models.Tree{
				Val: 8,
			},
		},
	}
	ans = 2
	if val := BinaryTreeVisibleNodesCounting(bTree); val != ans {
		t.Errorf("Tree=%v. Expect %v, but get %v", "(8, (2, (8, None, None), (7, None, None)), (6, None, None)) >> 2 (8,8)", ans, val)
	}
}
