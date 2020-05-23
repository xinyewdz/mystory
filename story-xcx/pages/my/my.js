// pages/my/my.js
const app = getApp();
Page({
  data: {
    user:{},
    isLogin:false
  },
  onLoad: function (options) {

  },
  getUserInfo(e){
    var userInfo = JSON.parse(e.detail.rawData);
    console.log(userInfo);
    var that = this;
    wx.login({
      success:function(res){
        if(res.code){
          var data = {
            code:res.code,
            nickName:userInfo.nickName,
            gender:userInfo.gender
          };
          app.postData("/login",data,function(respData){
            that.setData({
              isLogin:true,
              user:respData
            });
            app.setUser(respData);
          });
        }else{
          console.log("login fail!"+res.errMsg);
          wx.showToast({
            title:"登录失败:"+res.errMsg,
            icon:"none"
          });
        }
      }
    });
  }
})