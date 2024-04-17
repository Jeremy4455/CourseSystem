// 当前页码
let currentPage = 1;

// 每页显示的条目数
const itemsPerPage = 5;

// 获取分页按钮
const prevPageBtn = document.getElementById('prev-page-btn');
const nextPageBtn = document.getElementById('next-page-btn');
// 检查是否已经在最后一页
const totalRows = document.querySelectorAll('tbody tr').length;
const totalPages = Math.ceil(totalRows / itemsPerPage);

// 上一页按钮点击事件处理函数
prevPageBtn.addEventListener('click', () => {
    if (currentPage > 1) {
        currentPage--;
        updatePage();
    }
});

// 下一页按钮点击事件处理函数
nextPageBtn.addEventListener('click', () => {
    if (currentPage < totalPages) {
        currentPage++;
        updatePage();
    }
});

// 更新页面内容
function updatePage() {
    // 计算当前页的起始索引
    const startIndex = (currentPage - 1) * itemsPerPage;
    // 获取所有课程行
    const courseRows = document.querySelectorAll('tbody tr');

    // 显示当前页的课程行，隐藏其余的
    courseRows.forEach((row, index) => {
        if (index >= startIndex && index < startIndex + itemsPerPage) {
            row.style.display = 'table-row';
        } else {
            row.style.display = 'none';
        }
    });
    // 更新分页按钮状态
    prevPageBtn.disabled = currentPage === 1;
    // 检查是否已经在最后一页，并禁用下一页按钮
    nextPageBtn.disabled = currentPage === totalPages;
}

// 初始化页面
updatePage();
