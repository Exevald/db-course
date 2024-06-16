import './index.css';
import {renderBranchListPage} from "./view/pages/BranchList/BranchList";
import {renderCreateBranchPage} from "./view/pages/CreateBranch/CreateBranch";
import {renderBranchInfoPage} from "./view/pages/BranchInfo/BranchInfo";
import {renderCreateEmployeePage} from "./view/pages/CreateEmployee/CreateEmployee";

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
    case "/employee/create": {
        renderCreateEmployeePage()
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