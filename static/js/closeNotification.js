function closeNotification(notificationId) {
    var notification = document.getElementById(notificationId);
    if (notification) {
        notification.style.display = 'none';
    }
}
