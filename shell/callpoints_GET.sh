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
echo -e "  3) DEVICES GET: https://xallcloud.appspot.com/api/devices "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/devices"


echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  4) ASSIGNMENTS GET: https://xallcloud.appspot.com/api/assignments/callpoint/UID-1000-0000-0001 "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/assignments/callpoint/UID-1000-0000-0001"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  5) ACTIONS GET: https://xallcloud.appspot.com/api/actions "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/actions"




echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  6) EVENTS GET: https://xallcloud.appspot.com/api/events "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/events"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  7) EVENTS GET: https://xallcloud.appspot.com/api/events/callpoint/UID-1000-0000-0001 "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/events/callpoint/UID-1000-0000-0001"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  8) EVENTS GET: https://xallcloud.appspot.com/api/events/action/UID-ACT-0000-0001 "
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/events/action/UID-ACT-0000-0001"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------"
echo -e "  100) GET API VERSION : https://xallcloud.appspot.com/api/version"
echo -e "-----------------------------------------------------------------------------------------------------"



curl -i -H "Accept: application/json" "https://xallcloud.appspot.com/api/version"



echo -e "\n"
echo -e "-----------------------------------------------------------------------------------------------------\n"
