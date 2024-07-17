VERSION=$1

echo ""
echo "------------------- COMMIT FIRST --------------------"
echo "------------------- COMMIT FIRST --------------------"
echo "------------------- COMMIT FIRST --------------------"
echo ""
echo "------- CHECK VERSION!!! Must start with \"v\"--------"
echo ""
echo "git tag $1"
echo "git push origin $1"
echo "GOPROXY=proxy.golang.org go list -m github.com/Juniper/terraform-provider-mist@$1"
echo ""