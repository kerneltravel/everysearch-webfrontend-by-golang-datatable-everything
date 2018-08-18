
function callmefirst() {
    //取得输入框JQuery对象 
    var $searchInput = $('#keyword');
    //关闭浏览器提供给输入框的自动完成 
    $searchInput.attr('autocomplete', 'on');
    var $searchResult = $('#searchresult');

    $("#searchbutton").click(
        function () {
            $("#searchresult").DataTable({
                "searching": true,
                "destroy": true,
                //'url':'http://localhost:8080/search', //服务器的地址
                "ordering": true,
                "serverSide": false,

                columns: [
                    { data: 'name' },
                    { data: 'path' },
                    { data: 'size' },
                    { data: 'fdate' }
                ],
                ajax: {
                    type: 'GET',
                    dataType: 'json',
                    dataSrc: 'data',
                    //后台golang编译后运行的服务，端口是8080
                    url: 'http://localhost:8080/search',
                    data: { 'keyword': $searchInput.val() }, //参数 
                    async: true,
                    error: function () {
                        alert("无法连接到后台搜索服务，请检查是否已运行？");
                    }
                }
            });
        }
    );

    $searchInput.bind('keypress', function (event) {
        $("#searchresult").DataTable({
            "searching": true,
            "destroy": true,
            //'url':'http://localhost:8080/search', //服务器的地址
            "ordering": true,
            "serverSide": false,
            columns: [
                { data: 'name' },
                { data: 'path' },
                { data: 'size' },
                { data: 'fdate' }
            ],
            ajax: {
                type: 'GET',
                dataType: 'json',
                dataSrc: 'data',
                url: 'http://localhost:8080/search',
                data: { 'keyword': $searchInput.val() }, //参数 
                async: true,
                error: function () {
                    alert("无法连接到后台搜索服务，请检查是否已运行？");
                }
            }
        });
    });
}; 