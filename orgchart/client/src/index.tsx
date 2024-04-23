import './index.css';
import {renderBranchListPage} from "./view/pages/BranchList/BranchList";
import {renderCreateBranchPage} from "./view/pages/CreateBranch/CreateBranch";
import {renderBranchInfoPage} from "./view/pages/BranchInfo/BranchInfo";

const location = window.location.pathname
switch (location) {
    case "/": {
        renderBranchListPage()
        break
    }
    case "/branch/create": {
        renderCreateBranchPage()
        break
    }
    case "/branch": {
        renderBranchInfoPage()
        break
    }
    default: {
        renderBranchListPage()
        break
    }
}