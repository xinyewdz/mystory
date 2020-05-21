// pages/story/play.js
const app = getApp()

Page({
  data: {
    story: {
      name:"",
      url:"",
      image:""
    }
  },
  onLoad: function (option) {
      var id = option.id;
      console.log("story:"+id);
      var that = this;
      this.getStory(id,function(story){
        that.setData(
          {
            "story":story
          }
        )
      });
     
  },
  getStory:function(id,callback){
    wx.request({
      url: app.host+'/story?id='+id,
      success:function(resp){
        var respData = resp.data;
        console.log(resp.data);
        if(respData.code==200){
          callback(respData.data)
        }
      }
    })
  }
})
