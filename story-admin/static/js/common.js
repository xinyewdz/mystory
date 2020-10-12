//var host = "https://api.story.wenqiuqiu.com";
var host = "http://localhost:8060";

var app={
    setToken:function(token){
        window.sessionStorage.setItem("token",token);
    },
    getToken:function(){
        return window.sessionStorage.getItem("token");
    },
    setUser:function(){
        window.sessionStorage.setItem("user",user);
    },
    getUser:function(){
        return window.sessionStorage.getItem("user");
    },
    request:function(url,data,succ,err){
        let token = app.getToken()
        let h={
            token:token
        }
        Vue.http.post(host+url,data,{headers:h}).then(resp=>{
            if(resp.body.code=="200"){
                succ(resp.body.data);
            }else if(resp.body.code=="401"){
                window.location.href="index.html";
            }else{
                alert(resp.body.msg);
                err(resp)
            }
        });
    }
}