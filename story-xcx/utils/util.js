const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : '0' + n
}

function getPickerIdx(val,data){
  let idx = 0;
  if(val=='undefined'||val==""){
    return idx;
  }
  for(;idx<data.length;idx++){
    if (data[idx].key==val){
      break;
    }
  }
  return idx;
}

function getPickerVal(idx,data){
  if(idx>=data.length){
    idx = data.length-1;
  }
  return data[idx].key;
}

module.exports = {
  formatTime: formatTime,
  getPickerIdx:getPickerIdx,
  getPickerVal:getPickerVal
}