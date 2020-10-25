import { shallow } from "enzyme";
import React from "react";
import InputFolder from "./InputFolder"

import axios from 'axios';
import API from '../api/api';

jest.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

describe("Test suite for InputFolder", ()=> {
    it("should render without crash", ()=> {
        const wrapper = shallow(<InputFolder folderName={""}/>);
        expect(wrapper).toBeTruthy();
    })

    it("should render equal to the snapshot", () => {
        const wrapper = shallow(<InputFolder folderName={""} />);
        expect(wrapper).toBeTruthy();
        expect(wrapper.debug()).toMatchSnapshot();
    });
    it("should make a post request with the right info", () => {
        const resp = {data: [{name: 'Bob'}]};
        mockedAxios.post.mockImplementation(() => Promise.resolve(resp))
        const wrapper = shallow(<InputFolder folderName={""} />);
        wrapper.find("button").simulate("click");
        expect(axios.post).toHaveBeenCalledTimes(1)
        expect(axios.post).toHaveBeenCalledWith("http://localhost:5000/api/v1/folder/", "")
    });

})
