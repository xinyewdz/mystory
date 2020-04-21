// pages/story/play.js
const app = getApp()

Page({
  data: {
    story: {}
  },
  onLoad: function (option) {
      console.log(option.query)
      var id = "1";
      var story = this.getStory(id);
      this.setData(
        {
          "story":story
        }
      )
  },
  getStory:function(id){
    var story = {
      "name":"test",
      "image":"http://up.wenqiuqiu.com/%E6%95%91%E6%8A%A4%E8%BD%A6.png",
      "url":"http://up.wenqiuqiu.com/audio/%E6%95%91%E6%8A%A4%E8%BD%A6%E5%A3%B0%E9%9F%B3.mp3"
    };
    return story;
  }
})
