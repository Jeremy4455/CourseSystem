function closeNotification(id) {
  var notification = document.getElementById(id);
  notification.parentNode.removeChild(notification);
}

// 自动关闭通知框
setTimeout(function() {
  var notification = document.getElementById('success-notification');
  if (notification) {
    closeNotification('success-notification');
  }
}, 3000);

setTimeout(function() {
  var notification = document.getElementById('error-notification');
  if (notification) {
    closeNotification('error-notification');
  }
}, 3000);
