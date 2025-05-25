package data

import (
	"testing"

	v1 "spring-go-rpc/api/gen/go/system/service/v1"

	"github.com/bytedance/sonic"
)

// TestSonicJSONSerialization 测试sonic JSON序列化功能
func TestSonicJSONSerialization(t *testing.T) {
	// 创建测试用户数据
	user := &v1.UserInfo{
		Id:       1,
		Username: "测试用户",
		Email:    "test@example.com",
		Age:      25,
	}

	// 测试序列化
	jsonData, err := sonic.Marshal(user)
	if err != nil {
		t.Fatalf("sonic序列化失败: %v", err)
	}

	// 测试反序列化
	var deserializedUser v1.UserInfo
	err = sonic.Unmarshal(jsonData, &deserializedUser)
	if err != nil {
		t.Fatalf("sonic反序列化失败: %v", err)
	}

	// 验证数据一致性
	if deserializedUser.Id != user.Id {
		t.Errorf("ID不匹配: 期望 %d, 实际 %d", user.Id, deserializedUser.Id)
	}
	if deserializedUser.Username != user.Username {
		t.Errorf("用户名不匹配: 期望 %s, 实际 %s", user.Username, deserializedUser.Username)
	}
	if deserializedUser.Email != user.Email {
		t.Errorf("邮箱不匹配: 期望 %s, 实际 %s", user.Email, deserializedUser.Email)
	}
	if deserializedUser.Age != user.Age {
		t.Errorf("年龄不匹配: 期望 %d, 实际 %d", user.Age, deserializedUser.Age)
	}

	t.Logf("sonic JSON序列化测试通过，JSON数据: %s", string(jsonData))
}
