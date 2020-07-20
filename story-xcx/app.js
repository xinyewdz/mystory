//app.js
App({
  host:"https://api.story.wenqiuqiu.com",
  adm:null,
  //host:"http://localhost:8060",
  onLaunch:function(){
   this.adm = wx.getBackgroundAudioManager();
  },
  onHide:function(){
    
  },
  setUser:function(user){
    wx.setStorage({
      data: user,
      key: 'user',
    });
    this.setToken(user.token);
  },
  getUser:function(){
    return wx.getStorageSync('user')
  },
  getToken:function(){
    return wx.getStorageSync('token')
  },
  setToken:function(token){
    wx.setStorage({
      data: token,
      key: 'token',
    })
  },
  auth:function(){
    var token = this.getToken();
    console.log(token)
    return token!=""
  },
  postData:function(url,data,callback){
    var that = this;
    wx.request({
      url: that.host+url,
      method:"POST",
      header:{
        token:that.getToken()
      },
      data:data,
      complete:function(resp){
        var respData = resp.data;
        if(respData.code==200){
          callback(respData.data);
        }else{
          wx.showToast({
            title: respData.msg,
            icon:"none"
          })
        }
      }
    })
  }
})