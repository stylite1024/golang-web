function testApi (data) {
    return $axios({
        'url': '/api/v1/test',
        'method': 'get',
        data
    })
}