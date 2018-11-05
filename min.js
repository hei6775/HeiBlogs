<script type="text/javascript">
    function createXMLHttpRequest() {
            try {
                return new XMLHttpRequest();
            } catch (e) {
                try {
                    return new ActiveXObject("Msxml2.XMLHTTP");
                } catch (e) {
                    return new ActiveXObject("Microsoft.XMLHTTP");
}
}
}

        function send() {
            var xmlHttp = createXMLHttpRequest();
            xmlHttp.onreadystatechange = function() {
                if(xmlHttp.readyState == 4 && xmlHttp.status == 200) {
                    if(xmlHttp.responseText == "true") {
        document.getElementById("error").innerText = "用户名已被注册！";
    document.getElementById("error").textContent = "用户名已被注册！";
                    } else {
        document.getElementById("error").innerText = "";
    document.getElementById("error").textContent = "";
}
}
};
xmlHttp.open("POST", "/ajax_check/", true, "json");
xmlHttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
var username = document.getElementById("username").value;
xmlHttp.send("username=" + username);
}
</script>