package main

import "testing"

func TestEncryptStringBySSHRsaPublicKeyAscii(t *testing.T) {
	encryptStringBySSHRsaPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6pGxKXhoASEQIHKkmF1YafXXEZbsuxier0PqEntA0CqGO0CiirxRECEpSKfST8ImjKw8M13tAsyyFxO2L9YgSYtBgW4GtNWqhEHETnsxYW9fUFjvHe7qpnVBjo+GWC0cU4IsPL0qsghzgroaXSHd0NIpMU4iwp/0LUIzaGIi/fnu/uWj4CxdbNOhK63EVQNbBVEXtOt+XmbQda0ygLTjrYhJ9CIaMwbF8uZJmlKgfvQFbnCks9wwZCrKmSHlVuzXKwPy3+qv0XXwwND6nwAhUU8A481qgKUSVxRBI65IB6JCJ8oL5AYB7OxQiuKFrkgwJzNydnfXYHY0qKFGFmCBF", "hello, welcome")
}

func TestEncryptStringBySSHRsaPublicKeyMultibytes(t *testing.T) {
	encryptStringBySSHRsaPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC6pGxKXhoASEQIHKkmF1YafXXEZbsuxier0PqEntA0CqGO0CiirxRECEpSKfST8ImjKw8M13tAsyyFxO2L9YgSYtBgW4GtNWqhEHETnsxYW9fUFjvHe7qpnVBjo+GWC0cU4IsPL0qsghzgroaXSHd0NIpMU4iwp/0LUIzaGIi/fnu/uWj4CxdbNOhK63EVQNbBVEXtOt+XmbQda0ygLTjrYhJ9CIaMwbF8uZJmlKgfvQFbnCks9wwZCrKmSHlVuzXKwPy3+qv0XXwwND6nwAhUU8A481qgKUSVxRBI65IB6JCJ8oL5AYB7OxQiuKFrkgwJzNydnfXYHY0qKFGFmCBF", "筋肉ムキムキ")
}
