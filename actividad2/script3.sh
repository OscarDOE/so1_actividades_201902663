GITHUB_USER=""  
echo "Ingresa tu nombre de usuario de GitHub:"
read GITHUB_USER
response=$(curl -sL https://api.github.com/users/$GITHUB_USER) #Get con curl                                                                                                                                                                     
echo  "Hola $GITHUB_USER User ID: $(echo $response | jq -r '.id'). La cuentaa fue creada el: $(echo $response | jq -r '.created_at')"                                                                                                           
crr_path="/tmp/$(date +%Y-%m-%d)"                                                                                       
mkdir -p $crr_path                                                                                                      
echo $msg >> "$crr_path/saludos.log"    