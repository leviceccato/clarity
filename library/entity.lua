local entityId = 1

return function()
    local e = {}
    e.id = entityId
    entityId = entityId + 1
    e.components = {}

    e.addComponent = function(component)
        e.components[#e.components + 1] = component.type
        e[component.type] = component.data
    end

    return e
end