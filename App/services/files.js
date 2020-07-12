async function getFiles() {
    try {
        const response = await axios.get('http://localhost:8080/files/');
        return response.data.data;
    } catch (error) {
        console.log(error);
        throw error;
    }
}

module.exports = getFiles;