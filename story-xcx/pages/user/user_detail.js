// pages/user/user_detail.js
const util = require('../../utils/util.js');
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    user:{
      phone:"",
      type:"",
      gender:"",
      name:""
    },
    genderData:[{val:"男",key:"男"},{key:"女",val:"女"}],
    userTypeData:[{val:"管理员",key:"admin"},{key:"user",val:"普通用户"}]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    let id = options.id;
    let that = this;
    let url = "/user/detail";
    app.postData(url,{id:id},function(respData){
      that.setData({
        user:respData
      })
    });
  },
  typeChange:function(event){
    let type = event.detail.value;
    this.data.user.type= type;
  },
  genderChange:function(event){
    let gender = event.detail.value;
    this.data.user.gender = gender;
  },
  handleInputChange: function (e) {
    let name = e.currentTarget.dataset.name;
    let value = e.detail.value;
    this.data.user[name] = value;
},
  save:function(){
    let user = this.data.user;
    app.postData("/user/save",user,function(){
      let url = "/pages/user/user_list";
      wx.redirectTo({
        url: url
      })
    });
  }
})