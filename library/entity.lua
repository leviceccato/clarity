local entityId = 1

return function()
    local e = {}
    e.id = entityId
    entityId = entityId + 1
    e.components = {}

    e.addComponent = function(component)
        e[component.type] = component.data
    end

    return e
end