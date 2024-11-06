
function createElementa(
    tag,
    props,
    children
) {

    return {
        tag,
        props: props || {},
        children: children || []
    }
}

function renderElement(vNode) {

    if (typeof vNode === "string") {
        
        return document.createTextNode(vNode);
    }

    const el = document.createElement(vNode.tag);
    for (const [ key, value ] of Object.entries(vNode.props)) {
        el.setAttribute(key, value);
    }

    vNode.children.forEach(child => {
        el.appendChild(renderElement(child));
    });

    return el;
}

function diff(oldNode, newNode) {

    const patches = [];
    if (oldNode === undefined || newNode === undefined) {

        if (newNode !== undefined) {
            patches.push({ type: "ADD", newNode });
        } else if (oldNode !== undefined) {
            patches.push({ type: "REMOVE" });
        }
    }

    else if (typeof oldNode === "string" && typeof newNode === "string") {
        if (oldNode !== newNode) {
            patches.push({ type: "TEXT", text: newNode });
        }
    }

    else if (oldNode.tag !== newNode.tag) {
        patches.push({ type: "REPLACE", newNode });
    }

    else {
        const propPatches = [];
        for (const [ key, value ] of Object.entries(newNode.props)) {
            if (oldNode.props[key] !== value) {
                propPatches.push({ key, value });
            }
        }

        for (const key in oldNode.props) {
            if (!(key in newNode.props)) {
                propPatches.push({ key });
            }
        }

        if (propPatches.length > 0) {
            patches.push({ type: "PROPS", props: propPatches })
        }

        const childPatch = []
        const maxChildrenLength = Math.max(oldNode.children.length, newNode.children.length);
        for (let i = 0; i < maxChildrenLength; i++) {
            childPatch.push(diff(oldNode.children[i], newNode.children[i]));
        }
        patches.push({ type: "CHILDREN", children: childPatch });
    }

    return patches;
}

function patch(parent, patches, index = 0) {

    const el = parent.children[index];
    
    patches.forEach(_patch => {
        switch(_patch.type) {
            case "ADD":
                parent.appendChild(renderElement(_patch.newNode));
                break;
            case "REMOVE":
                parent.removeChild(el);
                break;
            case "TEXT":
                el.textContent = _patch.text;
                break;
            case "REPLACE":
                parent.replaceChild(renderElement(_patch.newNode), el);
                break;
            case "PROPS":
                _patch.props.forEach(({ key, value }) => {
                    if (value === undefined){
                        el.removeAttribute(key);
                    } else {
                        el.setAttribute(key, value);
                    }
                });
                break;
            case "CHILDREN":
                _patch.children.forEach((childPatch, i) => {
                    patch(el, childPatch, i);
                });
                break;
        }
    });
}

window.onload = function () {

    const runVirtualDOM = document.getElementById("runVirtualDOM");
    let oldDom = createElementa("div", {}, [
        createElementa("p", {}, ["Hello World!"]),
        createElementa("ol", {
            start: 0
        }, [
            createElementa("li", {
                class: "aaa",
            }, ["你好世界！"]),
        ]),
    ]);

    console.log("Create Element Node:");
    console.log(oldDom);

    console.log("Render to browser:");
    console.log(renderElement(oldDom));

    runVirtualDOM.appendChild(renderElement(oldDom));

    let newDom = createElementa("div", {}, [
        createElementa("p", {}, ["Hello World!"]),
        createElementa("ol", {
            start: 0
        }, [
            createElementa("li", {}, ["你好世界！"]),
            createElementa("li", {
                class: "aaa",
            }, ["你好世界！!"]),
        ]),
        createElementa("p", {
            style: "color: #adb",
        }, ["P tag"])
    ]);

    console.log("New Element Node:");
    console.log(newDom);

    let domPatch = diff(oldDom, newDom);
    console.log("Diffrent");
    console.log(domPatch);

    setTimeout(() => {
        patch(runVirtualDOM, domPatch);
    }, 5000);
}