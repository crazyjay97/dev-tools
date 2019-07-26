package {{packageName}}.{{moduleName}}.service.impl;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import java.util.List;
import java.util.Map;
import com.baomidou.mybatisplus.plugins.Page;
import com.baomidou.mybatisplus.service.impl.ServiceImpl;
import {{mainPath}}.common.utils.PageUtils;
import {{mainPath}}.common.utils.Query;

import {{packageName}}.{{moduleName}}.dao.{{className}}Dao;
import {{packageName}}.{{moduleName}}.entity.{{className}}Entity;
import {{packageName}}.{{moduleName}}.service.{{className}}Service;


/**
  * {{ table.TableComment }}
  *
  * @author {{ authorName }}
  * @email {{ emailAddress }}
  * @date {{ genTime }}
  */
@Service("{{ table.FileName }}Service")
public class {{className}}ServiceImpl extends ServiceImpl<{{className}}Dao, {{className}}Entity> implements {{className}}Service {

    @Override
    public PageUtils queryPage(Map<String, Object> params) {
        Page page = new Query(params).getPage();
        page.setRecords(baseMapper.page(page, params));
        return new PageUtils(page);
    }

    @Override
    public boolean update(${className}Entity entity) {
    //todo
        return super.updateById(entity);
    }

    @Override
    public boolean insert(${className}Entity entity) {
    //todo
        return super.insert(entity);
    }

    @Override
    public void deleteBatchIds(List<String> ids) {
        baseMapper.deleteBatchIds(ids);
     }
}
