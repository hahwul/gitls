echo "[+] Start init"
echo "[+] Prject name?"
read templatehahwul
echo "[+] Your project name is $templatehahwul"

# 1
echo "[+] Change .goreleaer.yml"
sed "s/TEMPLATE-HAHWUL/\$templatehahwul/" .goreleaser.yml | tee .goreleaser.yml

#2
echo "[+] Make 'go.mod'"
go mod init github.com/hahwul/$templatehahwul

#3 
echo "[+] Cobra init"
cobra init --pkg-name github.com/hahwul/$templatehahwul
cobra add version

#final
echo "[+] Delete init.sh"
git rm init.sh
echo "[+] Please check '.goreleaser.yml' for descriptions"
