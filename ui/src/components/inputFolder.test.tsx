import { shallow } from "enzyme";
import React from "react";
import InputFolder from "./InputFolder"

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
})
