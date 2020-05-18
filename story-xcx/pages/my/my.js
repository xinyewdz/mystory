// pages/my/my.js
const app = getApp()
Page({
  data: {
    isAdmin:true,
    login:false,
    adminName:"aaron",
    user:{}
  },
  onLoad: function (options) {
    //this.getAdminName();
  },
  getAdminName:function(){
    var that = this;
    wx.request({
      url: app.host+"/adminName",
      success:function(resp){
        var name = resp.data;
        that.setData(
          {
            adminName:name
          }
        )
      }
    })
  },
  getUserInfo(e){
    var userInfo = JSON.parse(e.detail.rawData);
    console.log(userInfo);
    var that = this;
    this.setData({
      user:userInfo,
      isAdmin: that.data.adminName==userInfo.nickName,
      login:true
    });
    
  }


})