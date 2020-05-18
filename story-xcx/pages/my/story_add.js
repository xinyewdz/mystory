// pages/my/my.js
const app = getApp()
const audio = wx.createInnerAudioContext();
Page({
  data: {
    "name":"",
    "image":"",
    "audioPath":"",
    "audioName":""
  },
  onLoad: function (options) {

  },
  saveStory:function(){
    wx.showLoading({
      "title":"上传中"
    });
    var that = this;
    let name = this.data.name;
    console.log("savestory,"+name);
    that.upload(that.data.audioPath,name,function(path){
      var audioPath = path;
      console.log("audio upload success");
      that.upload(that.data.image,name,function(path){
        console.log("image upload success");
        var imagePath = path;
        wx.request({
          url: app.host+"/save",
          method:"POST",
          data:{
            name:name,
            audio:audioPath,
            image:imagePath
          },
          success:function(resp){
            wx.hideLoading();
            const respData = resp.data;
            if(respData.code=="200"){
              wx.switchTab({
                url: '../index/index',
                success:function(){
                  var page = getCurrentPages().pop()
                  if(page==undefined||page==null) return;
                  page.onLoad();
                }
              })
            }
          }
        })
      });
    });
    
  },
  upload:function(path,name,callback){
    wx.uploadFile({
      filePath: path,
      name: "file",
      url: app.host+"/upload",
      formData:{
        "name":name
      },
      success:function(res,code){
        let resData = JSON.parse(res.data);
        let fp = resData.data;
        console.log("upload success.fp="+fp);
        callback(fp);
      },
      fail:function(){
        console.log("upload fail.path="+path);
      }
    })
  },
  nameEvent:function(event){
    var name = event.detail.value;
    console.log("name event."+name);
    this.setData({
      "name":name
    })
  },
  testEvent:function(){
    console.log("testevent")
  },
  chooseAudio:function(){
    var that = this;
    console.log("chooseAudio event")
    wx.chooseMessageFile({
        count:1,
        type:"all",
        success (res) {
          const tempFilePaths = res.tempFiles;
          that.setData({
            "audioPath":tempFilePaths[0]["path"],
            "audioName":tempFilePaths[0]["name"]
          })
        }
      })
  },
  chooseImg:function(){
    var that = this;
    console.log("chooseImg event")
    wx.chooseImage({
      success (res) {
        const tempFilePaths = res.tempFilePaths;
        that.setData({
          "image":tempFilePaths[0]
        })
      }
    })
  }

})