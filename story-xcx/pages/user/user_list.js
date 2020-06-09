// pages/my/user_list.js
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    userList:[]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
      this.getUserList();
  },
  getUserList:function(){
    let that = this;
    app.postData("/user/list",{},function(respData){
        that.setData({
          userList:respData
        })
    });
  },
  userDetail:function(event){
    let id = event.currentTarget.id;
    console.log("user "+id);
    let url = "/pages/user/user_detail?id="+id;
    wx.redirectTo({
      url: url
    })
  }
})