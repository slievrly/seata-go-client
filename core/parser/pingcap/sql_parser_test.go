/**
2 * @Author: Nico
3 * @Date: 2019/9/7 20:02
4 */
package pingcap

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewMySQLRecognizer(t *testing.T) {
	sql := "select name, age from tbl t where id = 1"

	mr, err := NewMySQLRecognizer(sql)
	assert.Equal(t, err, nil)
	assert.Equal(t, mr.StmtNode.Text(), sql)
	assert.Equal(t, mr.TableName, "tbl")
	assert.Equal(t, mr.TableAlias, "t")
}
