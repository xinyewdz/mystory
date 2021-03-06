// pages/my/my.js
const app = getApp();
Page({
  data: {
    user:{},
    isLogin:false
  },
  onShow: function (options) {
  },
  logout:function(){
    this.setData({
      user:{},
      isLogin:false
    });
    wx.removeStorage({
      key: 'user',
    });
    wx.removeStorage({
      key: 'token',
    })
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
            gender:userInfo.gender+"",
            avatarUrl:userInfo.avatarUrl,
            province:userInfo.province,
            city:userInfo.city
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