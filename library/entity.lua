local entityId = 1

return function()
    local e = {}
    e.id = entityId
    entityId = entityId + 1

    e.addComponent = function(component)
        e[component.type] = component
    end

    return e
end