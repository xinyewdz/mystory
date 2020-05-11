// pages/story/play.js
const app = getApp()

Page({
  data: {
    story: {}
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
        console.log(JSON.stringify(respData));
        if(respData.code=="200"){
          callback(respData.data)
        }
      }
    })
  },
  remove:function(event){
      var id = event.currentTarget.dataset.id;
      console.log("remove "+id);
      wx.request({
        url: app.host+'/remove?id='+id,
        success:function(resp){
          var respData = resp.data
          wx.switchTab({
            url: '../index/index',
            success:function(){
              var page = getCurrentPages().pop()
              if(page==undefined||page==null) return;
              page.onLoad();
            }
          })
        }
      })
  }
})
