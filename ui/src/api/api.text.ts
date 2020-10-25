import axios from 'axios';
import API from './api';

jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

describe("Test suite for API", ()=> {

    it('should fetch users', () => {
        const api = new API()
        const resp = {data: [{name: 'Bob'}]};
        mockedAxios.post.mockImplementation(() => Promise.resolve(resp))

        api.getFolderInfo("xx").then(data => expect(data).toEqual(resp));
    });
})
