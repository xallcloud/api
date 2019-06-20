echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  1) GET API VERSION : https://xallcloud.appspot.com/api/version"
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/version"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  2)  CALLPOINTS GET: https://xallcloud.appspot.com/api/callpoints "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/callpoints"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  DEVICES GET: https://xallcloud.appspot.com/api/devices "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/devices"


echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  ASSIGNMENTS GET: https://xallcloud.appspot.com/api/assignments/callpoint/UID-1000-0000-0001 "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/assignments/callpoint/UID-1000-0000-0001"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  ACTIONS GET: https://xallcloud.appspot.com/api/actions "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/actions"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  100) GET API VERSION : https://xallcloud.appspot.com/api/version"
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/version"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------\n"
